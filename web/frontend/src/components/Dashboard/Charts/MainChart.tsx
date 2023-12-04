import { Line } from 'react-chartjs-2';
import { CodingTimeOutput } from "./../../../api/types/codingTime"


import {
    Chart as ChartJS,
    LineElement,
    CategoryScale,
    LinearScale,
    PointElement,
    Filler,
    ChartOptions
} from 'chart.js';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Filler
);

export function MainChart(props: CodingTimeOutput) {
    const data = {
        labels: props.dates,
        datasets: [
            {
                text: props.duration,
                data: props.times,
                borderColor: '#ef6461',
                backgroundColor: 'transparent',
                pointBorderColor: '#ef6461',
                borderRadius: 4,
                tension: 0.3,
                fill: {
                    target: "origin",
                    above: "rgba(255, 0, 0, 0.3)"
                }
            },
        ],
    }

    const options : ChartOptions = {
        responsive: true,
        plugins: {
            legend: {
                display: false
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
        scales: {
            x: {
                grid: {
                    display: false,
                }
            },
            y: {
                grid: {
                    display: false,
                }
            }
        }
    };

    return (
        <>
            {props.dates.length === 0 && props.times.length === 0
                ? <p>No Data Found</p>
                : <Line options={options} data={data} />}
        </>
    )
}
