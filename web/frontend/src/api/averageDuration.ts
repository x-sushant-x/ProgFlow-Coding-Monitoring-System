import axios from 'axios'
import { AverageTime } from './types/averageTime'

export class AverageTimeAPI {
    static async getAverageTime(username : string) {
        const resp = (await axios.get(`http://localhost:8080/analytics/average-time?username=${username}`)).data['data']

        const output : AverageTime = {averageTime: 0, todayTime: 0, todayDurationText: '', averageDurationText: ''}

        output.todayTime = resp['todayTime']
        output.todayDurationText = resp['todayDurationText']
        output.averageTime = resp['averageTime']
        output.averageDurationText = resp['averageDurationText']

        console.log('Average Time: ' + output)

        return output
    }
}