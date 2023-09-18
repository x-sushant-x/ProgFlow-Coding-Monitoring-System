import React, { useState } from 'react';
import { FaHome, FaHandsHelping } from 'react-icons/fa';
import { IoPeopleSharp } from 'react-icons/io5';
import { GoGoal } from 'react-icons/go';
import { ImStatsBars } from 'react-icons/im'
import { useNavigate } from 'react-router-dom';

interface MenuItem {
  icon: React.ElementType;
  text: string;
  link: string
}

const menuItems: MenuItem[] = [
  { icon: FaHome, text: 'Home', link: '/dashboard' },
  { icon: ImStatsBars, text: 'Stats', link: '/stats' },
  { icon: FaHandsHelping, text: 'Your Team', link: '/team' },
  { icon: GoGoal, text: 'Goals', link: '/goals' },
  { icon: IoPeopleSharp, text: 'Leaderboard', link: '/leaderboard' },
  // { icon: BsShareFill, text: 'Share', link: '/share'},

];

export default function SideBar() {
  const [activeButton, setActiveButton] = useState<number | null>(0)

  const navigate = useNavigate()
  const navigateTo = (index: number) => navigate(menuItems[index]['link'])

  const HandleButtonClick = (index: number) => {
    navigateTo(index)
    setActiveButton(index);
  };



  return (
    <div className='sidebar bg-white border-2 h-screen'>

      <p className='mt-[1rem] ml-6 text-2xl font-bold text-gray-800'>ProgFlow</p>

      <div className='menu mt-28 ml-6'>
        {menuItems.map((item, index) => {
          const Icon = item.icon;
          return (
            <div
              key={index}
              className={`bg-gray-300 w-36 h-8 rounded flex items-center mt-3 ${activeButton === index ? 'bg-blue-500' : ''
                }`}
              onClick={() => HandleButtonClick(index)}
            >
              <Icon
                style={{
                  color: activeButton === index ? 'white' : 'grey',
                  fontSize: '17px',
                }}
                className='ml-2'
              />
              <p
                className={`ml-3 text-sm ${activeButton === index ? 'text-white' : 'text-gray-500'
                  }`}
              >
                {item.text}
              </p>
            </div>
          );
        })}
      </div>
    </div>
  );
}
