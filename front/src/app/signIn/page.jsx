'use client'

import { useEffect, useState } from 'react';
import { supabase} from '../supabase'

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
    <div className="h-screen w-screen flex justify-center items-center">

    <Card className="w-[60%]">
      <CardHeader>
        <CardTitle>サインイン</CardTitle>
        <CardDescription>Eメールとパスワードでログインしましょう。</CardDescription>
        </CardHeader>
        <CardContent>

        <form>
              <Label htmlFor="email">Email</Label>
              <Input id="email" placeholder="Email of your project" onChange={(e) => setEmail(e.target.value)} />
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="Password">Password</Label>
              <Input id="password" type='password' onChange={(e) => setPassword(e.target.value)}/>
              <Label htmlFor="name">Name</Label>
              <Input id="name" placeholder="Name of your project" />
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


