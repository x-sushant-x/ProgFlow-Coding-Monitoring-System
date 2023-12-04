export interface CodingTimeData {
    time: string
    date: string
    duration: string
}

export interface CodingTimeOutput {
    times: string[]
    dates: string[]
    duration: string[]
}

export interface CodingStatisticsOutput {
    today : string
    thisWeek : string
    thisMonth : string
    allTime : string
}