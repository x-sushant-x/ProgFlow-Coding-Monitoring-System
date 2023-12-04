import { useAuth0 } from '@auth0/auth0-react'
import { Chart as ChartJS, ArcElement, Tooltip, Legend, ChartOptions } from 'chart.js'
import { useEffect, useState } from 'react'
import { Doughnut } from 'react-chartjs-2'
import { ProjectAPI } from '../../api/projectDuration'
import { useMyContext } from '../../contexts/Context'

ChartJS.register(ArcElement, Tooltip, Legend)

export function ProjectTime() {
    const { user } = useAuth0()
    const [projectsName, setProjectsName] = useState<string[]>([])
    const [time, setTime] = useState<number[]>([])
    const [timeDuration, setTimeDuration] = useState<string[]>([])

    const { days } = useMyContext()



    useEffect(() => {
        (async () => {
            const projectTime = await ProjectAPI.getProjectTime(user?.nickname ?? '', days)
            setProjectsName(projectTime.projectName)
            setTime(projectTime.totalTime)
            setTimeDuration(projectTime.totalTimeDuration)
        })()
    }, [user, days])

    const data = {
        labels: projectsName,
        datasets: [
            {
                time: timeDuration,
                label: 'Time Spent: ',
                data: time,
                backgroundColor: [
                    'rgba(54, 162, 235, 1)',
                    'rgba(153, 102, 255, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(255, 127, 14, 1)',
                    'rgba(255, 99, 132, 1)',
                    'rgba(255, 206, 86, 1)',
                    'rgba(255, 159, 64, 1)',
                ],
                borderWidth: 1,
                spacing: 2,
            },
        ]
    }

    const options: ChartOptions = {
        plugins: {
            legend: {
                position: 'right',
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
                        const time = data.datasets[datasetIndex].time[dataIndex]
                        return `${time}`
                    }
                }
            },
        },
    };

    return (
        projectsName.length === 0 ?
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
