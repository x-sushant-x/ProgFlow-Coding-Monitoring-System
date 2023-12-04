import { useEffect, useState } from "react"
import axios from 'axios'
import { Leaderboard } from "../../api/types/leaderboard"

interface LeaderboardProps {
    rank: number
    username: string
    languages: string
    total_duration: number
    week_duration_text: string
    daily_duration_text: string
}

export default function LeaderboardPage() {
    const [leaderboard, setLeaderboard] = useState<Leaderboard[]>()

    useEffect(() => {
        (async () => {
            try {
                const leaderboardResp = await (await axios.get('http://localhost:8080/analytics/leaderboard')).data['data']
                setLeaderboard(leaderboardResp)
            } catch (err) {
                console.log(err)
            }
        })()
    }, [])

    const headingTextStyle = 'text-sm font-bold text-slate-900 ml-4'

    return (
        <>
            <div className='mx-4 my-8 bg-white showdow rounded shadow h-12 w-[77rem] flex  items-center'>
                <p className={headingTextStyle}>Rank</p>
                <p className={headingTextStyle + ' ml-8'}>Users</p>
                <p className={headingTextStyle + ' ml-48'}>Week</p>
                <p className={headingTextStyle + ' ml-32'}>Daily</p>
                <p className={headingTextStyle + ' ml-32'}>Languages</p>
            </div>


            {
                leaderboard?.map(el => {
                    
                    const userData : LeaderboardProps = {
                        rank: el.rank,
                        username: el.username,
                        languages: el.languages,
                        total_duration: el.total_duration,
                        daily_duration_text: el.daily_duration_text,
                        week_duration_text: el.week_duration_text
                    }
                    console.log(userData)
                    return LeaderboardUserCard(userData)
                })
            }

        </>
    )
}

function LeaderboardUserCard(props : LeaderboardProps) {
    
    const textStyle = 'text-sm text-slate-900 ml-4'

    return (
        <>
        <div className={'mx-4 bg-white my-2 showdow rounded shadow h-10 w-[77rem] flex items-center '} key={props.rank}>
                <div className={textStyle + " w-[4rem] h-[1rem] ml-[1rem]"}>{props.rank}</div>
                <div className={textStyle + " w-[13rem] h-[1rem] ml-[0.4rem]"}>{props.username}</div>
                <div className={textStyle + " w-[10rem] h-[1rem] ml-[1.2rem]"}> {props.week_duration_text}</div>
                <div className={textStyle + " w-[10rem] h-[1rem] ml-[0.7rem]"}>{props.daily_duration_text}</div>
                <div className={textStyle + " w-[38rem] h-[1rem] ml-[0.7rem]"}>{props.languages}</div>
            </div>
        </>
    )
}