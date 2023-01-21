import { useState } from 'react'
import './App.css'

function App() {


  return (
    <div className="App">
      <div className="container h-full flex flex-col ">
        <div className="top flex justify-center items-center h-16 border-solid border-2 text-blue-600">
          top
        </div>
        <div className="main flex justify-center items-center h-96 border-solid border-2">
          main
        </div>
        <div className="bottom p-5 flex justify-between items-center h-16 border-solid border-2">
          <div className="btn1">抽</div>
          <div className="btn2">城</div>
          <div className="btn3">副</div>
          <div className="btn4">野</div>
          <div className="btn5">包</div>
        </div>
      </div>
    </div>
  )
}

export default App
