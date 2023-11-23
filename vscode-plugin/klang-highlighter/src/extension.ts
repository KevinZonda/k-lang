'use strict';

import * as net from 'net';
import { Trace } from 'vscode-jsonrpc';
import { workspace, ExtensionContext, OutputChannel, window } from 'vscode';

import { LanguageClient, LanguageClientOptions } from 'vscode-languageclient/node';

let client: LanguageClient;

let enableLsp : boolean = true;

let outputChannel: OutputChannel;

export function activate(context: ExtensionContext) {
    outputChannel = window.createOutputChannel("Klang LSP");
    // outputChannel.show(true);
    // outputChannel.appendLine('activate');
    if (enableLsp) {
        startLsp()
    }
}

function startLsp() {
    const socket = net.connect(11451, '127.0.0.1');
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