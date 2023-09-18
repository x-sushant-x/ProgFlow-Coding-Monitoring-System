import { useNavigate } from 'react-router-dom'
import { useAuth0 } from '@auth0/auth0-react'
import { useEffect } from 'react'


export default function HeaderHome() {
    const navigate = useNavigate()
    const { loginWithRedirect } = useAuth0()
    const { user, isAuthenticated } = useAuth0()

    useEffect(() => {
        if (isAuthenticated) {
            navigate('/dashboard')
        }
    }, [isAuthenticated, navigate])


    const login = async () => {
        loginWithRedirect()
    }

    const navigateToDashboard = () => {
        navigate('/dashboard')
        console.log(user)
    }



    return (
        <>
            <div className='flex justify-between mx-[8rem] pt-6 items-center'>
                <p className='text-2xl font-bold text-gray-800'>ProgFlow</p>



                <ul className='flex space-x-10 text-slate-700'>
                    <li>Home</li>
                    <li>About</li>
                    <li>Contact</li>
                </ul>



                {isAuthenticated ?

                    <div className='flex w-28 h-10 bg-slate-900 rounded-md text-white items-center justify-center cursor-pointer'>
                        <p onClick={navigateToDashboard}>
                            See Stats
                        </p>
                    </div> :
                    <div className='flex w-28 h-10 bg-slate-900 rounded-md text-white items-center justify-center cursor-pointer'>
                        <p onClick={login}>
                            Sign Up
                        </p>
                    </div>
                }
            </div>
        </>
    )
}