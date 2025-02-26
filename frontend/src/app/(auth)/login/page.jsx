import React from 'react'

export default function page() {
    //logic
  return (
    //view
    <div>
      <h1>zellcode</h1>
      <form>
        <input id='username' required placeholder='enter email or username'> </input>
        <input type='password' required placeholder='enter your password' />
        <button type = 'submit'>Login</button>
      </form>
    </div>
  )
}
