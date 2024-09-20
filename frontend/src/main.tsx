import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import Payment from './Payment.tsx'
import InternetBanking from './InternetBanking.tsx'
import Wallet from './E-Wallet.tsx'
import Paypal from './Paypal.tsx'
import Cards from './Cards.tsx'
import './index.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    {/* <Payment /> */}
    {/* <InternetBanking /> */}
    {/* <Paypal /> */}
    <Cards />
    {/* <Wallet /> */}
  </StrictMode>,
)
