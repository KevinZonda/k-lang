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

let lsp : any;
let isLspReady = false;


function pullUpLspServer(addr: string) {
    let cmd = klangCfg.get<string>("runner");
    if (cmd === undefined || cmd === "" || cmd === null) {
        cmd = "klang"
    }
    lspChannel.appendLine("LSP Path: " + cmd)
    const spawn = require("child_process").spawn;
    lsp = spawn(
        cmd, ["lsp", addr],
        {
            detatched: true,
        }
    );

    lsp.stdout.on("data", (data) => {
        lspChannel.append(data.toString());
        lspChannel.appendLine("Lsp ready (STDOUT)");
        isLspReady = true;
    });

    lsp.stderr.on("data", (data) => {
        lspChannel.append(data.toString());
        lspChannel.appendLine("Lsp ready (STDERR)");
        isLspReady = true;
    });
    lsp.on("close", (code) => {
        isLspReady = true;
        runChannel.appendLine("Lsp closed: " + code);
    });
}

function strToHostPort(addr: string) : [string, number] {
    let host = ""
    let port = 11451;
    let i = addr.indexOf(":");
    if (i >= 0) {
        host = addr.substring(0, i);
        port = parseInt(addr.substring(i + 1));
    }
    return [host, port];
}

async function startLsp() {
    let lspListenAddr = klangCfg.get<string>("lspListenAddress");
    if (lspListenAddr === undefined || lspListenAddr === "" || lspListenAddr === null) {
        lspListenAddr = "127.0.0.1:11451"
    }
    pullUpLspServer(lspListenAddr);
    lspChannel.appendLine("LSP Listen Address: " + lspListenAddr);

    while (isLspReady === false) {
        await new Promise(resolve => setTimeout(resolve, 1000));
        lspChannel.appendLine("Waiting for LSP to be ready...");
    }
    await new Promise(resolve => setTimeout(resolve, 5000));
    lspChannel.appendLine("LSP is ready");
    const host = strToHostPort(lspListenAddr);
    const socket = net.connect(host[1], host[0]);
    socket.on('connect', () => {
        lspChannel.appendLine("LSC connected");
    });
    socket.on('close', () => {
        lspChannel.appendLine("LSC closed");
    });
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
    if (lsp) {
        lsp.kill();
    }
    return client.stop();
}