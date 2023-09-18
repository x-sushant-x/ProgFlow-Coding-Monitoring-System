import { Chart as ChartJS, ArcElement, Tooltip, Legend, ChartOptions } from 'chart.js'
import { useEffect, useState } from 'react'
import { Doughnut } from 'react-chartjs-2'
import { LanguageAPI } from '../../api/language-api';
import { useAuth0 } from '@auth0/auth0-react';
import { languageLogoColors } from '../../utils/languageColors';
import { useMyContext } from '../../contexts/Context';

ChartJS.register(ArcElement, Tooltip, Legend)

export function LanguageChart() {
    const { user } = useAuth0()

    const [languages, setLanguages] = useState<string[]>([])
    const [textDuration, setTextDuration] = useState<string[]>([])
    const [time, setTime] = useState<number[]>([])

    const { days } = useMyContext()


    useEffect(() => {
        (async () => {
            try {
                const languageTimeResp = await LanguageAPI.getLanguageTime(user?.nickname ?? '', days)
                setLanguages(languageTimeResp.languages)
                setTextDuration(languageTimeResp.durationText)
                setTime(languageTimeResp.times)
            } catch (e) {
                console.log(e)
            }
        })()
    }, [user?.nickname, days])


    const data = {
        // labels: ['GoLang', 'JavaScript', 'C++', 'Make', 'Python'],
        labels: languages,
        datasets: [
            {
                time: textDuration,
                label: 'Time Spent: ',
                // data: [45, 20, 25, 2.5, 2.5],
                data: time,
                // backgroundColor: [
                //     'rgba(255, 99, 132, 1)',
                //     'rgba(54, 162, 235, 1)',
                //     'rgba(255, 206, 86, 1)',
                //     'rgba(75, 192, 192, 1)',
                //     'rgba(153, 102, 255, 1)',
                //     'rgba(255, 159, 64, 1)',
                //     'rgba(255, 127, 14, 1)'
                // ],

                backgroundColor: languages.map(lang => {
                    return languageLogoColors[lang]
                }),
                
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
                        return `Time: ${time}`
                    }
                }
            },
        },
    };

    return (

        languages.length === 0 ?
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
