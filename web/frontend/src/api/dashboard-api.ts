import axios from "axios"
import { CodingStatisticsOutput, CodingTimeData, CodingTimeOutput } from "./types/codingTime"

export class DashboardAPI {
    static async getCodingTime(username: string, days: string) {
        const baseURL = `http://localhost:8080/analytics/coding-time?username=${username}&days=${days}`

        const resp = await axios.get(baseURL)

        const timeData: CodingTimeData[] = resp.data['data']

        const output: CodingTimeOutput = { times: [], dates: [], duration: [] }

        timeData.map(el => {
            output.times.push(el.time)
            output.dates.push(el.date)
            output.duration.push(el.duration)
        })

        output.dates.reverse()
        output.times.reverse()
        output.duration.reverse()

        console.log('Coding Time' + output.duration)

        return output
    }

    static async getCodingStatistics(username: string) {
        const apiURL = `http://localhost:8080/analytics/coding-statistics?username=${username}`
        const resp = await axios.get(apiURL)

        const statisticsData : CodingStatisticsOutput = resp.data['data']

        console.log(statisticsData)
        
        return statisticsData
    }
}
