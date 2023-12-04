import axios from "axios"
import { LanguageTimeData, LanguageTimeOutput } from "./types/languageTime"

export class LanguageAPI {
    static async getLanguageTime(username: string, days: string) {
        const URL = `http://localhost:8080/analytics/language-time?username=${username}&days=${days}`

        const resp = await axios.get(URL)

        const languageTimeData : LanguageTimeData[] = resp.data['data']
        const output : LanguageTimeOutput = { languages: [], times : [], durationText: [] }

        languageTimeData.map(el => {
            output.languages.push(el.languageName)
            output.times.push(el.totalDuration)
            output.durationText.push(el.durationText)
        })

        console.log('Language Time: ', output)

        return output
    }
}