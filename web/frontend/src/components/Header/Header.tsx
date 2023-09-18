import './Header.css';
import React, { ReactNode } from 'react';
import { MdLogout } from 'react-icons/md';
import { RxAvatar } from 'react-icons/rx';
import Box from '@mui/material/Box'
import MenuItem from '@mui/material/MenuItem'
import FormControl from '@mui/material/FormControl'
import Select, { SelectChangeEvent } from '@mui/material/Select'
import { useMyContext } from '../../contexts/Context';


const Header: React.FC = () => {
    const { days, setDays } = useMyContext()

    const handleChange = (event: SelectChangeEvent) => {
        setDays(event.target.value as string)
    }

    console.log('Value: ' + days)


    return (
        <div className="header bg-white border-b-2 flex justify-between h-14">
            <div className='flex items-center ml-12 '>
                <h3 className='mr-6 text-slate-900'>Select Time: </h3>
                <Box sx={{ minWidth: 100 }}>
                    <FormControl fullWidth>
                        <Select
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            value={days}
                            onChange={handleChange}
                        >
                            <MenuItem value={1}>Today</MenuItem>
                            <MenuItem value={7}>Last 7 Days</MenuItem>
                            <MenuItem value={30}>Last Month</MenuItem>
                            <MenuItem value={365}>Last Year</MenuItem>
                            <MenuItem value={3650}>All Time</MenuItem>
                        </Select>
                    </FormControl>
                </Box>
            </div>

            <div className='icons flex items-center space-x-6'>
                <IconWrapper>
                    <MdLogout style={{ fontSize: '25px' }} />
                </IconWrapper>
                <IconWrapper>
                    <RxAvatar style={{ fontSize: '25px' }} className='mr-[1rem]' />
                </IconWrapper>
            </div>
        </div>
    )
}

interface IconWrapperProps {
    children: ReactNode;
}

const IconWrapper: React.FC<IconWrapperProps> = ({ children }) => {
    return <div className='icon text-slate-800'>{children}</div>;
}

export default Header;
