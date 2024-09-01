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

  const onSubmit = async(e) => {
      e.preventDefault();
      try{
        const { error:signUpError } = await supabase.auth.signUp({
          email: email,
          password: password,
        })
        if (signUpError) {
          throw signUpError;
        }
      alert('登録完了メールを確認してください');
      }catch(error){
        alert('エラーが発生しました');
      }
    };

  return (
    <div className="w-screen h-screen flex flex-row items-center justify-center">

    <Card className="w-[60%]">
      <CardHeader>
        <CardTitle>サインアップ</CardTitle>
        <CardDescription>Eメールとパスワードでサインアップしましょう。</CardDescription>
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
        <Button onClick={onSubmit}>サインアップ</Button>
        <a href='/signIn'>サインインへ</a>
      </CardFooter>
    </Card>
    </div>
  )
}


