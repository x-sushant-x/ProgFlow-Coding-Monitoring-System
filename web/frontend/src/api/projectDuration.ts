import axios from "axios"
import { ProjectTime, ProjectTimeOutput } from "./types/projectTime"

export class ProjectAPI {
    static async getProjectTime(username: string, days: string) {
        const URL = `http://localhost:8080/analytics/project-time?username=${username}&days=${days}`

        const resp = await axios.get(URL)

        const apiResp : ProjectTime[] = resp.data['data']
        const output : ProjectTimeOutput = {
            projectName: [],
            totalTime: [],
            totalTimeDuration: []
        }

        apiResp.map(el => {
            output.projectName.push(el.projectName)
            output.totalTime.push(el.totalTime)
            output.totalTimeDuration.push(el.totalTimeDuration)
        })

        console.log('Project Data' + output)
        return output
    }
}