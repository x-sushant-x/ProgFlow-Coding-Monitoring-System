/* eslint-disable @typescript-eslint/naming-convention */
import * as vscode from 'vscode'
import http = require('http')
import axios from 'axios'
require('dotenv').config()

export class APIHandler {
    async addProject(ctx: vscode.ExtensionContext, name: string) {
        const apiKey: string = ctx.globalState.get('progflow.apiKey') ?? ""
        const baseURL = 'http://localhost:8080/project/add'

        const postData = JSON.stringify({
            name: name
        })
        const options = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Content-Length': postData.length,
                'x-api-key': apiKey
            }
        }

        const request = http.request(baseURL, options, (response) => {
            let data = ''

            response.on('data', (chunk) => {
                data += chunk.toString()
            })

            response.on('end', () => {
                console.log('Response:', data)
            });
        })

        request.on('error', (error) => {
            console.log('Error:', error)
        })

        request.write(postData)
        request.end()
    }


    async updateCodingActivity(ctx: vscode.ExtensionContext, name: string, startTime: string) {
        try {
            const baseURL = 'http://localhost:8080/coding-activity'
            const apiKey: string = ctx.globalState.get('progflow.apiKey') ?? ""

            const headers = {
                'Content-Type': 'application/json',
                'x-api-key': apiKey
            }

            let response = await axios.post(baseURL, { projectName: name, startTime: startTime }, { headers })
            console.log(response.data)
        } catch (err) {
            console.log(err)
        }
    }

    async updateLanguageActivity(ctx: vscode.ExtensionContext, name: string, languageName: string, startTime: string) {
        try {
            const baseURL = 'http://localhost:8080/language-activity'
            const apiKey: string = ctx.globalState.get('progflow.apiKey') ?? ""

            const headers = {
                'Content-Type': 'application/json',
                'x-api-key': apiKey
            }

            let response = await axios.post(baseURL, { projectName: name, languageName: languageName, startTime: startTime }, { headers })
            console.log(response.data)
        } catch (err) {
            console.log(err)
        }
    }
}