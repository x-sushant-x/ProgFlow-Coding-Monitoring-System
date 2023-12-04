import { Grid } from '@mui/material'
import { LanguageChart } from './LanguageStatsChart'
import { AverageComparison } from './AverageComparison'
import { ProjectTime } from './ProjectTime'
import ComingSoon from './ComingSoon'

export default function Stats() {
    const displayBox = 'border-[1px] rounded-md shadow-md flex items-center justify-center'
    return (
        <>
            <div className='p-12'>
                <Grid container spacing={4}>
                    <Grid item md={6}>
                        <div className={displayBox} style={{ aspectRatio: '6/3', backgroundColor: '#ffffff', overflow: 'hidden' }}>
                            <LanguageChart />
                        </div>
                        <p></p>
                    </Grid>
                    <Grid item md={6}>
                        <div className={displayBox} style={{ padding: '0.5rem', aspectRatio: '6/3', backgroundColor: '#ffffff', overflow: 'hidden' }}>
                            <AverageComparison />
                        </div>
                    </Grid>


                    <Grid item md={6}>
                        <div className={displayBox} style={{ padding: '0.5rem', aspectRatio: '6/3', backgroundColor: '#ffffff', overflow: 'hidden' }}>
                            <ProjectTime />
                        </div>
                    </Grid>
                    <Grid item md={6}>
                        <div className={displayBox} style={{ padding: '0.5rem', aspectRatio: '6/3', backgroundColor: '#ffffff', overflow: 'hidden' }}>
                            <ComingSoon />
                        </div>
                    </Grid>
                </Grid>
            </div>
        </>
    )
}