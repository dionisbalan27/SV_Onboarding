import {
  AlipayCircleOutlined,
  LockOutlined,
  MobileOutlined,
  TaobaoCircleOutlined,
  UserOutlined,
  WeiboCircleOutlined,
} from '@ant-design/icons';
import { Alert, message, Tabs } from 'antd';
import React, { useState } from 'react';
import { ProFormCaptcha, ProFormCheckbox, ProFormText, LoginForm } from '@ant-design/pro-form';
import { useIntl, history, FormattedMessage, SelectLang, useModel } from 'umi';
import Footer from '@/components/Footer';
import { login } from '@/services/ant-design-pro/api';
import { getFakeCaptcha } from '@/services/ant-design-pro/login';
import styles from './index.less';

const LoginMessage = ({ content }) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

const Login = () => {
  const [userLoginState, setUserLoginState] = useState({});
  const [type, setType] = useState('account');
  const { initialState, setInitialState } = useModel('@@initialState');
  const intl = useIntl();

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();

    if (userInfo) {
      await setInitialState((s) => ({ ...s, currentUser: userInfo }));
    }
  };

  const handleSubmit = async (values) => {
    //const { personalNumber, password } = values;
    try {
      const msg = await login(values);
      if (msg.status === 'ok') {
        console.log(msg.data.role);
        localStorage.setItem('token', msg.data.token);
        localStorage.setItem('username', msg.data.name);
        localStorage.setItem('role', msg.data.role);
        const defaultLoginSuccessMessage = intl.formatMessage({
          id: 'pages.login.success',
          defaultMessage: '登录成功！',
        });
        message.success(defaultLoginSuccessMessage);
        await fetchUserInfo();
        /** 此方法会跳转到 redirect 参数所在的位置 */

        // if (!history) return;
        // const { query } = history.location;
        // const { redirect } = query;
        if (msg.data.role == 'signer') {
          console.log('success');
          history.push('/signer');
          return;
        }
        if (msg.data.role == 'admin') {
          history.push('/admin');
          return;
        }
        if (msg.data.role === 'maker') {
          history.push('/');
          return;
        }
        if (msg.data.role === 'checker') {
          history.push('/checker');
          return;
        } else {
          history.push(redirect || '/viewer');
          return;
        }
      }

      console.log(msg); // 如果失败去设置用户错误信息

      setUserLoginState(msg);
    } catch (error) {
      const defaultLoginFailureMessage = intl.formatMessage({
        id: 'pages.login.failure',
        defaultMessage: '登录失败，请重试！',
      });
      message.error(defaultLoginFailureMessage);
    }
  };

  // try {
  //   // 登录
  //   const msg = await login({ ...values, type });

  //   if (msg.status === 'ok') {
  //     const defaultLoginSuccessMessage = intl.formatMessage({
  //       id: 'pages.login.success',
  //       defaultMessage: '登录成功！',
  //     });
  //     message.success(defaultLoginSuccessMessage);
  //     const test = await fetchUserInfo();
  //     /** 此方法会跳转到 redirect 参数所在的位置 */
  //     console.log(test);

  //     if (!history) return;
  //     const { query } = history.location;
  //     const { redirect } = query;
  //     history.push(redirect || '/');
  //     return;
  //   }

  //   console.log(msg); // 如果失败去设置用户错误信息

  //   setUserLoginState(msg);
  // } catch (error) {
  //   const defaultLoginFailureMessage = intl.formatMessage({
  //     id: 'pages.login.failure',
  //     defaultMessage: '登录失败，请重试！',
  //   });
  //   message.error(defaultLoginFailureMessage);
  // }

  const { status, type: loginType } = userLoginState;
  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang />}
      </div>
      <div className={styles.content}>
        <LoginForm
          logo={<img alt="logo" src="/logo.svg" />}
          title="Ant Design"
          subTitle={intl.formatMessage({
            id: 'pages.layouts.userLayout.title',
          })}
          initialValues={{
            autoLogin: true,
          }}
          onFinish={async (values) => {
            await handleSubmit(values);
          }}
        >
          <Tabs activeKey={type} onChange={setType}>
            <Tabs.TabPane
              key="account"
              tab={intl.formatMessage({
                id: 'pages.login.accountLogin.tab',
                defaultMessage: '账户密码登录',
              })}
            />
          </Tabs>

          {status === 'error' && loginType === 'account' && (
            <LoginMessage
              content={intl.formatMessage({
                id: 'pages.login.accountLogin.errorMessage',
                defaultMessage: '账户或密码错误(admin/ant.design)',
              })}
            />
          )}
          {type === 'account' && (
            <>
              <ProFormText
                name="personalNumber"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.login.username.placeholder',
                  defaultMessage: '用户名: admin or user',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.login.username.required"
                        defaultMessage="请输入用户名!"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText.Password
                name="password"
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.login.password.placeholder',
                  defaultMessage: '密码: ant.design',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.login.password.required"
                        defaultMessage="请输入密码！"
                      />
                    ),
                  },
                ]}
              />
            </>
          )}
          <div
            style={{
              marginBottom: 24,
            }}
          >
            <ProFormCheckbox noStyle name="autoLogin">
              <FormattedMessage id="pages.login.rememberMe" defaultMessage="自动登录" />
            </ProFormCheckbox>
            <a
              style={{
                float: 'right',
              }}
              href="/user/register"
            >
              <FormattedMessage
                id="pages.login.forgotPasswords"
                defaultMessage="Don't have account ? Register"
              />
            </a>
          </div>
        </LoginForm>
      </div>
      <Footer />
    </div>
  );
};

export default Login;
