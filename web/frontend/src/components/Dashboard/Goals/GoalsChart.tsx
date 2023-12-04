import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';

ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
);

// eslint-disable-next-line react-refresh/only-export-components
export const options = {
    responsive: true,
    plugins: {
        legend: {
            display: false
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

const labels = ['23 August', '24 August', '25 August', '26 August', '27 August', '28 August', '29 August'];

// eslint-disable-next-line react-refresh/only-export-components
export const data = {
    labels,
    datasets: [
        {
            data: [12, 56, 34, 76, 41, 55, 60],
            backgroundColor: [
                '#E23030',
                '#3cb043'
            ],
            borderRadius: 5
        }
    ],
};

export function GoalsChart() {
    return <Bar options={options} data={data} />;
}
