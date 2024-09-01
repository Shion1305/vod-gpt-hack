'use client'

import { useEffect, useState } from 'react';
import { supabase} from '../supabase'
import { useRouter } from "next/navigation";
import * as React from "react"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

export default function SignIn() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const router = useRouter()

  const onSubmit = async () => {
    try {
      await supabase.auth.signInWithPassword({email, password})
      router.push('/')
    } catch (e) {
      console.log(e) 
    }
  };

  return (
    <div className="w-screen h-screen flex flex-row items-center justify-center">

    <Card className="w-[60%]">
      <CardHeader>
        <CardTitle>サインイン</CardTitle>
        <CardDescription>Eメールとパスワードでログインしましょう。</CardDescription>
      </CardHeader>
      <CardContent>
        <form>
          <div className="grid w-full items-center gap-4">
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="email">Email</Label>
              <Input id="email" placeholder="Email of your project" onChange={(e) => setEmail(e.target.value)} />
            </div>
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="Password">Password</Label>
              <Input id="password" type='password' onChange={(e) => setPassword(e.target.value)}/>
            </div>
          </div>
        </form>
      </CardContent>
      <CardFooter className="flex justify-between">
        <Button onClick={onSubmit}>サインイン</Button>
        <a href='/signUp'>サインアップへ</a>
      </CardFooter>
    </Card>
    </div>
  )
}


