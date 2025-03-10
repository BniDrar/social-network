import React from 'react'

export default function page() {
    //logic
  return (
    //view
    <div>
      <h1>zellcode</h1>
      <form>
        <input id='username' placeholder='enter email or username' required/>
        <input type='password' placeholder='enter your password' required/> 
        <button type = 'submit'>Login</button>
      </form>
    </div>
  )
}
