import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Dashboard from './components/Dashboard/Dashboard'
import Layout from './layout/Layout'
import Teams from './components/Teams/Teams'
import GoalsPage from './components/Dashboard/Goals/GoalsPage'
import Stats from './components/Stats/Stats'
import HomePage from './components/HomePage/HomePage'
import LeaderboardPage from './components/Leaderboard/LeaderboardPage'

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route index element = {<HomePage/>} />
          <Route path='/' element={<Layout />}>
            <Route path='/dashboard'  element={<Dashboard />} />
            <Route path='/stats' element={<Stats />} />
            <Route path='/team' element= {<Teams/>}/>
            <Route path='/goals' element={<GoalsPage/>} />
            <Route path='/leaderboard' element={<LeaderboardPage/>} />
            <Route path='/share' element={<div>Share Page</div>} />
          </Route>
        </Routes>
      </Router>
    </>
  )
}

export default App