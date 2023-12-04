export interface LanguageTimeData {
    languageName: string
    totalDuration: number
    durationText: string
}

export interface LanguageTimeOutput {
    languages: string[]
    times : number[]
    durationText: string[] 
}