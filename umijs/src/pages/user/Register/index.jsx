import { Alert, message, Tabs } from 'antd';
import { useState } from 'react';
import { ProFormText, LoginForm } from '@ant-design/pro-form';
import { useIntl, FormattedMessage, SelectLang, history } from 'umi';
import Footer from '@/components/Footer';
import { register } from '@/services/ant-design-pro/api';
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

const Register = () => {
  const [type, setType] = useState('account');
  const intl = useIntl();

  const handleSubmit = async (values) => {
    const { personalNumber, password, name, email } = values;

    try {
      const res = await register({ personalNumber, password, name, email });
      if (res.status === 'ok') {
        const defaultLoginSuccessMessage = intl.formatMessage({
          id: 'pages.register.success',
          defaultMessage: 'Register success',
        });
        message.success(defaultLoginSuccessMessage);
        history.push('/user/login');
        return;
      }
    } catch (error) {
      const defaultLoginFailureMessage = intl.formatMessage({
        id: 'pages.register.failure',
        defaultMessage: 'Register Failed',
      });
      message.error(defaultLoginFailureMessage);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang />}
      </div>
      <div className={styles.content}>
        <LoginForm
          submitter={{
            render: (props, doms) => {
              console.log(props);
              return (
                <div style={{ width: '100%' }}>
                  <button
                    type="button "
                    key="submit "
                    style={{
                      width: '100%',
                      backgroundColor: '#1890ff',
                      border: 'none',
                      padding: 8,
                      color: 'white',
                    }}
                  >
                    Register
                  </button>
                </div>
              );
            },
          }}
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
                id: 'pages.register.accountLogin.tab',
                defaultMessage: 'Register Account',
              })}
            />
          </Tabs>

          {status === 'error' && loginType === 'account' && (
            <LoginMessage
              content={intl.formatMessage({
                id: 'pages.register.accountLogin.errorMessage',
                defaultMessage: '账户或密码错误(admin/ant.design)',
              })}
            />
          )}

          {type === 'account' && (
            <>
              <ProFormText
                name="name"
                placeholder={intl.formatMessage({
                  id: 'pages.register.username.placeholder',
                  defaultMessage: 'Name',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="Name is required"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText
                name="email"
                placeholder={intl.formatMessage({
                  id: 'pages.register.username.placeholder',
                  defaultMessage: 'Email',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="Email is required"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText
                name="personalNumber"
                placeholder={intl.formatMessage({
                  id: 'pages.register.username.placeholder',
                  defaultMessage: 'Personal Number',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="Personal number required"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText.Password
                name="password"
                placeholder={intl.formatMessage({
                  id: 'pages.register.password.placeholder',
                  defaultMessage: 'Password',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.password.required"
                        defaultMessage="Password required"
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
              paddingBottom: 20,
            }}
          >
            <a
              style={{
                float: 'right',
              }}
              onClick={() => history.push('/user/login')}
            >
              <FormattedMessage
                id="pages.login.redirectRegister"
                defaultMessage="Already have account, Login now !"
              />
            </a>
          </div>
        </LoginForm>
      </div>
      <Footer />
    </div>
  );
};

export default Register;
