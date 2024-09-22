import axios from 'axios';

export const PaymentService = {
  getPayments: async () => {
    const response = await axios.get('http://localhost:8020/api/payments');
    return response.data; // คืนค่าข้อมูลการชำระเงิน
  },
  createPayment: async (paymentData: any) => {
    const response = await axios.post('http://localhost:8020/api/payment', paymentData);
    return response.data; // คืนค่าผลลัพธ์การชำระเงิน
  },

};
