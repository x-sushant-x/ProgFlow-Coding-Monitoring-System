import * as vscode from 'vscode'
import { ProgFlow } from './ProgFlow';
import { Utils } from './utils/getDetails';
import path = require('path');
import { TimeUtils } from './utils/time';

export function activate(ctx: vscode.ExtensionContext) {
	let progFlow = new ProgFlow()
	let utils = new Utils();

	progFlow.startSession(ctx)

	/*----------------------------------------Commands -----------------------------------------------*/ 

	let newSession = vscode.commands.registerCommand('progflow.startSession', () => {
		progFlow.startSession(ctx)
	})
	ctx.subscriptions.push(newSession)


	let closeSession = vscode.commands.registerCommand('progflow.closeSession', () => {
		progFlow.ds.addData('Close Session')
	})

	vscode.commands.registerCommand('progflow.setAPIKey', () => {
		progFlow.setAPIKey(ctx)
	})

	ctx.subscriptions.pop()
	ctx.subscriptions.push(closeSession)


	/*----------------------- Setting and Clearning Start Time For Languages ----------------------------------*/ 

	// Setting start time for coding languages if their files are already open
	const openedLanguages = utils.getLanguages()
	openedLanguages.map(language => {
		ctx.globalState.update(`progflow.languageStartTime-${language}`, TimeUtils.getFormattedTime())
		console.log(`Setted Start Time For ${language} : ${TimeUtils.getFormattedTime()}`)
	})

	// Setting start time for language if a new language file is created and it's time is not available
	ctx.subscriptions.push(vscode.workspace.onDidOpenTextDocument((doc) => {
		const language = path.extname(doc.fileName)
		ctx.globalState.update(`progflow.languageStartTime-${language}`, TimeUtils.getFormattedTime())
		console.log(`Setted Start Time For ${language} : ${TimeUtils.getFormattedTime()}`)
	}))

	// When a single file is closed it's language coding time will reset throug this event
	ctx.subscriptions.push(vscode.workspace.onDidCloseTextDocument((doc) => {
		const closedDoc = doc.fileName
		const language = path.extname(closedDoc)

		ctx.globalState.update(`progflow.languageStartTime-${language}`, undefined)
		console.log(`Cleared start time for ${language}`)
	}))
}

export function deactivate(ctx: vscode.ExtensionContext) {
	let progFlow = new ProgFlow()
	progFlow.closeSession()		
}