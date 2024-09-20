import React from 'react';
import { Layout } from 'antd';
import { Divider, List, Typography } from 'antd';
import { Button, Flex } from 'antd';
import { Card } from 'antd';
import { Form, Input } from 'antd';
import imageSrc from './assets/Screenshot 2024-09-19 023621.png';


const { Header, Footer, Content } = Layout;

const headerStyle: React.CSSProperties = {
    display: 'flex',           // Use flexbox for alignment
    alignItems: 'center',       // Center vertically
    color: '#fff',
    height: 64,
    paddingInline: 48,
    lineHeight: '64px',
    backgroundColor: '#69ABC1',
    width: '100vw',             // Full width of the viewport
  };
  const contentStyle: React.CSSProperties = {
    minHeight: 'calc(100vh - 128px)',  // Full height minus header and footer
    lineHeight: '120px',
    backgroundColor: '#FFFFFF',
    width: '100vw',  // Full width of the viewport
  };
  
  const footerStyle: React.CSSProperties = {
    backgroundColor: '#FFFFFF',
    height: 64,
    width: '100vw',  // Full width of the viewport
  };
  
  const layoutStyle: React.CSSProperties = {
    borderRadius: 0,  // No border radius for full screen
    overflow: 'hidden',
    width: '100vw',  // Full width
    height: '100vh',  // Full height
  };
  const buttonStyle: React.CSSProperties = {
    textAlign: 'center',
    color: '#5F212E',
    paddingInline: 24,
    backgroundColor: '#FFFFFF',
    
  };

  const buttonclickStyle: React.CSSProperties = {
    textAlign: 'center',
    color: '#FFFFFF',
    backgroundColor: '#69ABC1',
    margin: '0'
  };

  const buttoncodeStyle: React.CSSProperties = {
    textAlign: 'center',
    color: '#FFFFFF',
    backgroundColor: '#69ABC1',
     marginLeft: '10px'
  };

  const listStyle: React.CSSProperties = { 
    padding: '20px 40px',
    width: '50%',           
    marginLeft: '5%',
          
  };

  const inputContainerStyle: React.CSSProperties = {
    display: 'flex',
    alignItems: 'center',  // จัดให้ทั้ง Input และ Button อยู่ตรงกลางแนวแกน Y
  };

  const headerContainerStyle: React.CSSProperties = {
    display: 'flex',
    alignItems: 'center',  // จัดให้ทั้ง Input และ Button อยู่ตรงกลางแนวแกน Y
  };

  const cardStyle: React.CSSProperties = {
    width: '30%',
    marginRight: '5%',
    marginTop: '-275px',  // เพิ่ม margin-top เพื่อให้ขนานกับ List
  };
  
  const data = [
    'Internet Banking',
    'Cards',
    'E-Wallet',
    'KBank',
  ];
  

const Payment: React.FC = () => (
  <Layout style={layoutStyle} >
      <Header style={headerStyle}>
        <div className='container'>
            <div className='topbar'>
              <div style={headerContainerStyle}>
                  <img src={imageSrc} alt="description" style={{ width: '10%', height: '10%',marginRight: '5%' }} />
                  <Flex gap="small" wrap>
                    <Button type="primary" style={buttonStyle}>Home</Button>
                    <Button type="primary" style={buttonStyle}>Fight</Button>
                    <Button type="primary" style={buttonStyle}>Benefits</Button>
                    <Button type="primary" style={buttonStyle}>Help Center</Button>
                  </Flex>
                  </div>
             </div>
        </div>
      </Header>
      <Content style={contentStyle}>
      <>
        <Divider orientation="left"></Divider>
        <List
            bordered
            dataSource={data}
            style={listStyle}
            renderItem={(item) => (
            <List.Item >
                <Typography.Text > {item}</Typography.Text>
                <Button type="primary" style={buttonclickStyle}>click</Button>
            </List.Item>
            )}
        />
        <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
          <Card title="Payment" bordered={false} style={cardStyle}>
                 
              <Form.Item
                layout="vertical"
                label="Voucher/Promo Code"
                name="Voucher/Promo Code"
                rules={[{ required: true }]}
                labelCol={{ span: 24 }}
                wrapperCol={{ span: 24 }}
              >
                <div style={inputContainerStyle}>
                  <Input style={{ flex: 1 }} />  {/* ใช้ flex: 1 เพื่อให้ Input ขยายเต็มพื้นที่ */}
                  <Button type="primary" style={buttoncodeStyle}>click</Button>
                </div>
                <br></br>
                Flight
                <br></br>
                <br></br>
                Price Detail
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                Total
                <br></br>
                <br></br>
                <br></br>
                PaymentStatus
              </Form.Item>
              
          </Card>
        </div>
        </>
      </Content>
      <Footer style={footerStyle}></Footer>
    </Layout>
);

export default Payment;