interface TopCodingTimeTextPrefs {
    title: string
    time: string
}

export default function TopCodingTimeText(props : TopCodingTimeTextPrefs) {
    return (
        <>
            <p className='text-white text-xs ml-2'>{props.title}</p>
            <p className='text-white text-md ml-2 mt-1'> {props.time.split(' ')[0]} {props.time.split(' ')[1]}</p>
            <p className='text-white text-md ml-2'> {props.time.split(' ')[2]} {props.time.split(' ')[3]}</p>
        </>
    )
}