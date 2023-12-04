import { Chart as ChartJS, ArcElement, Tooltip, Legend, ChartOptions } from 'chart.js'
import { useEffect, useState } from 'react'
import { Doughnut } from 'react-chartjs-2'
import { AverageTime } from '../../api/types/averageTime'
import { AverageTimeAPI } from '../../api/averageDuration'
import { useAuth0 } from '@auth0/auth0-react'
import { useLocation } from 'react-router-dom'

ChartJS.register(ArcElement, Tooltip, Legend)

export function AverageComparison() {
    const { user } = useAuth0()

    const location = useLocation()

    const { pathname } = location

    const [todayTime, setTodayTime] = useState(0)
    const [averageTime, setAverageTime] = useState(0)
    const [todayDurationText, setTodayDurationText] = useState('No Data')
    const [averageDurationText, setAverageDurationText] = useState('No Data')

    let averageTimeOutput: AverageTime = {
        averageDurationText: '',
        averageTime: 0,
        todayDurationText: '',
        todayTime: 0,
    }

    useEffect(() => {
        (async function () {
            try {
                // eslint-disable-next-line react-hooks/exhaustive-deps
                averageTimeOutput = await AverageTimeAPI.getAverageTime(user?.nickname ?? '')

                setTodayTime(averageTimeOutput.todayTime)
                setAverageTime(averageTimeOutput.averageTime)
                setTodayDurationText(averageTimeOutput.todayDurationText)
                setAverageDurationText(averageTimeOutput.averageDurationText)

            } catch (error) {
                console.log(error);
            }
        })()
    }, []);



    const data = {
        labels: ['Today', 'Average'],
        datasets: [
            {
                text: [todayDurationText, averageDurationText],
                label: 'Time Spent: ',
                data: [todayTime, averageTime],
                backgroundColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgb(3, 121, 212)'
                ],
                borderWidth: 1,
                spacing: 2,
            },
        ]
    }

    const options: ChartOptions = {
        plugins: {
            legend: {
                position: pathname === '/dashboard' ? 'top' : 'right',
                align: 'center',
                labels: {
                    boxWidth: 10,
                },
            },
            tooltip: {
                callbacks: {
                    label: (context): string => {
                        const datasetIndex = context.datasetIndex ?? 0
                        const dataIndex = context.dataIndex ?? 0
                        const time = data.datasets[datasetIndex].text[dataIndex]
                        return `${time}`
                    }
                }
            },
        },
    };

    return (
        todayTime === 0 ?
            <div className='text-center'>
                <p>No Data Found</p>
                <p>See how to fix this <a href='' className='text-blue-700'>here</a></p>
            </div>
            :
            <div>
                <Doughnut data={data} options={options} />
            </div>
    );
}
