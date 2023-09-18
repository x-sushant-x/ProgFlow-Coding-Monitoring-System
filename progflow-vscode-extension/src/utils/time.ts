export class TimeUtils {
    static getFormattedTime(): string {
        const date = new Date()
        let hour = date.getHours()
        let mins = date.getMinutes()

        return hour + ':' + mins
    }
}