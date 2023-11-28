'use strict';

import * as net from 'net';
import { Trace } from 'vscode-jsonrpc';
import { workspace, ExtensionContext, OutputChannel, window } from 'vscode';
import { basename, dirname, extname, join } from "path";
import * as vscode from "vscode";

import { LanguageClient, LanguageClientOptions } from 'vscode-languageclient/node';

let client: LanguageClient;

let enableLsp : boolean = true;

let lspChannel: OutputChannel;
let runChannel: OutputChannel;
let klangCfg : vscode.WorkspaceConfiguration

function runCode(file : string) {
    const spawn = require("child_process").spawn;
    let cmd = klangCfg.get<string>("runner");
    if (cmd === undefined || cmd === "" || cmd === null) {
        cmd = "klang"
    }
    runChannel.appendLine("Runner: " + cmd);
    runChannel.appendLine("File  : " + file)
    runChannel.appendLine("---------------------------------")
    let p = spawn(
        cmd, [file],
        {
            cwd:  dirname(file),
            shell: true
        }
    );

    p.stdout.on("data", (data) => {
        runChannel.append(data.toString());
    });

    p.stderr.on("data", (data) => {
        runChannel.append(data.toString());
    });

    p.on("close", (code) => {
        runChannel.appendLine("---------------------------------")
        runChannel.appendLine("Return: " + code);
    });
}

export function activate(context: ExtensionContext) {
    klangCfg = vscode.workspace.getConfiguration("klang")
    lspChannel = window.createOutputChannel("Klang LSP");
    runChannel = window.createOutputChannel("Klang Run");

    const run = vscode.commands.registerCommand("klang.run", (fileUri: vscode.Uri) => {
        runChannel.clear();
        runChannel.show(true);
        runCode(fileUri.fsPath);
    });
    context.subscriptions.push(run);
    // outputChannel.show(true);
    // outputChannel.appendLine('activate');
    if (enableLsp) {
        startLsp()
    }
}

function startLsp() {
    let lspListenAddr = klangCfg.get<string>("lspListenAddress");
    if (lspListenAddr === undefined || lspListenAddr === "" || lspListenAddr === null) {
        lspListenAddr = "127.0.0.1:11451"
    }
    lspChannel.appendLine("LSP Listen Address: " + lspListenAddr);
    const socket = net.connect(lspListenAddr);
    const serverOptions = () => {
        return Promise.resolve({
            writer: socket,
            reader: socket
        });
    };
    
    const clientOptions: LanguageClientOptions = {
        documentSelector: ['klang'],
        synchronize: {
            fileEvents: workspace.createFileSystemWatcher('**/*.k')
        }
    };
    
    client = new LanguageClient(
        'kLanguageServer',
        'Klang Language Server',
        serverOptions,
        clientOptions
    );

    client.setTrace(Trace.Verbose);
    client.start()
}

export function deactivate() : Thenable<void> | undefined {
    if (!client) {
		return undefined;
	}
    return client.stop();
}