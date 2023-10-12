import { Outlet } from 'react-router-dom'

function App() {
  return (
    <div className='app dark:bg-zinc-800 light:bg-white'>
        <Outlet />
    </div>
  )
}

export default App
