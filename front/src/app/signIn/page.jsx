'use client'

// import styles from '../styles/Home.module.css'
// 現時点で使わないものもあるが今後のことを考えて入れておく
import { Col, Container, Form, FormGroup, Input, Label, Row, Button } from "reactstrap";
import { useEffect, useState } from 'react';
// supabase
import { supabase } from '../utils/supabase';

// ヘッダーコンポーネントをインポート
import  Header from '../components/Header';
// useRouterをインポート
import { useRouter } from 'next/router';

export default function Register() {
  // useStateでユーザーが入力したメールアドレスとパスワードをemailとpasswordに格納する
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();

  // supabaseのユーザー登録の関数
  const doLogin =  async () => {
    // supabaseで用意されているユーザー登録の関数
    const { data, error } = await supabase.auth.signInWithPassword({ email, password })
    if (error) throw new Error(error.message)
    console.log(data)
    // ログインを反映させるためにリロードさせる
    router.reload()
  }

  return (
    // Home.module.cssでcardクラスに適用されているCCSを、このdivタグに適用する
    <div className={styles.card}>
      <h1>ログイン</h1>
      <Header/>
      <div>
        <Form>
            <FormGroup>
              <Label>
                メールアドレス：
              </Label>
              <Input
                type="email"
                name="email"
                style={{ height: 50, fontSize: "1.2rem" }}
                // onChangeでユーザーが入力した値を取得し、その値をemailに入れる
                onChange={(e) => setEmail(e.target.value)}
              />
            </FormGroup>
            <FormGroup>
              <Label>
                パスワード：
              </Label>
              <Input
                type="password"
                name="password"
                style={{ height: 50, fontSize: "1.2rem" }}
                // onChangeでユーザーが入力した値を取得し、その値をpasswordに入れる
                onChange={(e) => setPassword(e.target.value)}
              />
            </FormGroup>
            <Button
                style={{ width: 220 }}
                color="primary"
                // 登録ボタンがクリックされたとき関数が実行されるようにする
                onClick={()=>{
                  doLogin();
                }}
              >
              ログイン
            </Button>
        </Form>
      </div>
    </div>
  )
}
