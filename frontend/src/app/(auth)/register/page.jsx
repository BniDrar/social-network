"use client";
import { useState } from "react";
import { InputField } from "@/components/Input/input.jsx";

// name
// nickname
// email
// password
// confirmPassword:
// first
// last
// dateOfBirth:
// wantCode

export default function RegisterPage() {
  return (
    <div>
      <h2>Register</h2>
      <form>
        <InputField
          label="Name"
          type="text"
          name="name"
          value={""}
          onChange={()=>{}}
          required
        />
        <InputField
          label="Nickname"
          type="text"
          name="nickname"
          value={""}
          onChange={()=>{}}
        />
        <InputField
          label="Email"
          type="email"
          name="email"
          value={""}
          onChange={()=>{}}
          required
        />
        <InputField
          label="First Name"
          type="text"
          name="first"
          value={""}
          onChange={()=>{}}
          required
        />
        <InputField
          label="Last Name"
          type="text"
          name="last"
          value={""}
          onChange={()=>{}}
          required
        />
        <InputField
          label="Date of Birth"
          type="date"
          name="dateOfBirth"
          value={""}
          onChange={()=>{}}
          required
        />
        <InputField
          label="Password"
          type="password"
          name="password"
          value={""}
          onChange={()=>{}}
          required
        />
        <InputField
          label="Confirm Password"
          type="password"
          name="confirmPassword"
          value={""}
          onChange={()=>{console.log('change confirm password')}}
          required
        />
      </form>
    </div>
  );
}
