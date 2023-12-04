import { Grid } from '@mui/material'
import './Dashboard.css'
import { MainChart } from './Charts/MainChart'
import TopCodingTimeText from './TopCodingTimeText/TopCodingTimeText'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { useAuth0 } from '@auth0/auth0-react'
import { DashboardAPI } from '../../api/dashboard-api'
import { CodingStatisticsOutput, CodingTimeOutput } from '../../api/types/codingTime'
import { AverageComparison } from '../Stats/AverageComparison'
import Share from './Share/Share'

export default function Dashboard() {
    const displayBox = 'border-[1px] rounded-md shadow-md'

    const { user } = useAuth0()
    const [dates, setDates] = useState<string[]>([])
    const [time, setTime] = useState<string[]>([])
    const [duration, setDuration] = useState<string[]>([])

    const [todayTime, setTodayTime] = useState('No Data');
    const [thisWeekTime, setThisWeekTime] = useState('No Data');
    const [thisMonthTime, setThisMonthTime] = useState('No Data');
    const [allTime, setAllTime] = useState('No Data');



    let codingTime: CodingTimeOutput = { times: [], dates: [], duration: [] }
    let stats : CodingStatisticsOutput = {
        today: 'No Data', thisWeek : 'No Data', thisMonth : 'No Data', allTime: 'No Data'
    };

    useEffect(() => {
        (async function () {
            try {
                await axios.post('http://localhost:8080/user', {
                    name: user?.name,
                    username: user?.nickname,
                    email: user?.email,
                    photo: user?.picture
                });

                // eslint-disable-next-line react-hooks/exhaustive-deps
                codingTime = await DashboardAPI.getCodingTime(user?.nickname ?? '', '7')
                setDates(codingTime.dates)
                setTime(codingTime.times)
                setDuration(codingTime.duration)

                // eslint-disable-next-line react-hooks/exhaustive-deps
                stats = await DashboardAPI.getCodingStatistics(user?.nickname ?? '')
                setTodayTime(stats.today)
                setThisWeekTime(stats.thisWeek)
                setThisMonthTime(stats.thisMonth)
                setAllTime(stats.allTime)

                console.log(thisMonthTime)

            } catch (error) {
                console.log(error);
            }
        })()
    }, []);



    return (
        <>
            <div style={{ padding: '3rem' }}>
                <Grid container>
                    <Grid item xs={12} md={9}>
                        <Grid container spacing={2}>
                            <Grid item xs={12} sm={6} md={2.5}>
                                <div className={displayBox} style={{ padding: '0.5rem', aspectRatio: '6/3', backgroundColor: '#F3737E', borderColor: '#F3737E' }}>
                                    <TopCodingTimeText title='Today' time = {todayTime} />
                                </div>
                            </Grid>

                            <Grid item xs={12} sm={6} md={2.5}>
                                <div className={displayBox} style={{ padding: '0.5rem', aspectRatio: '6/3', backgroundColor: '#18B797', borderColor: '#18B797' }}>
                                    <TopCodingTimeText title='This Week' time = {thisWeekTime} />
                                </div>
                            </Grid>


                            <Grid item xs={12} sm={6} md={2.5}>
                                <div className={displayBox} style={{ padding: '0.5rem', aspectRatio: '6/3', backgroundColor: '#3786EB', borderColor: '#3786EB' }}>
                                    <TopCodingTimeText title='This Month' time = {thisMonthTime} />
                                </div>
                            </Grid>


                            <Grid item xs={12} sm={6} md={2.5}>
                                <div className={displayBox} style={{ padding: '0.5rem', aspectRatio: '6/3', backgroundColor: '#F9AA4B', borderColor: '#F9AA4B' }}>
                                    <TopCodingTimeText title='All Time' time = {allTime} />
                                </div>
                            </Grid>

                            <Grid item xs={12} sm={6} md={10}>
                                <div className={displayBox} style={{ padding: '20px', textAlign: 'center', aspectRatio: '6/3', backgroundColor: '#ffffff' }}>
                                    <MainChart dates={dates} times={time} duration={duration} />
                                </div>
                            </Grid>
                        </Grid>
                    </Grid>

                    <Grid item md={3}>
                        <div className={displayBox} style={{ padding: '20px', textAlign: 'center', aspectRatio: '6/3', backgroundColor: '#ffffff' }}>
                            <AverageComparison/>
                        </div>
                        <div className='border-[1px] rounded-md shadow-md mt-7 h-[10.3rem]' style={{ padding: '20px', textAlign: 'center', backgroundColor: '#ffffff' }}><Share/></div>
                    </Grid>
                </Grid>

                <div className={displayBox} style={{ padding: '20px', textAlign: 'center', height: '5rem', marginTop: '2rem', backgroundColor: '#ffffff' }}>Ads</div>

            </div>

        </>
    )
}