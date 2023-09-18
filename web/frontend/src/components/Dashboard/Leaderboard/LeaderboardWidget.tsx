import { MdKeyboardArrowRight } from 'react-icons/md'

export default function LeaderboardWidget() {

    return (
        <>
            {/* User Component */}
            <div className='flex items-center justify-between mt-4'>
                <div className='flex items-center'>
                    <div className='w-6 h-6 bg-gray-600 rounded-full'></div>
                    <p className='ml-2 text-sm'>Sushant Dhiman</p>
                </div>
                <MdKeyboardArrowRight style={{ fontSize: '23px', color: 'gray' }} />
            </div>
        </>
    )
}