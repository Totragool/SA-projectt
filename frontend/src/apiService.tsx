import axios from 'axios';
import { Member } from './member.interface';
import { Booking } from './booking.interface';
import { Benefits } from './benefits.interface';
import { Payment } from './payment.interface';


// ตั้งค่า URL ของ API
const API_URL = 'http://localhost:8020'; // เปลี่ยนตาม URL ของ API ของคุณ

// Member API Service
export const MemberService = {
    getMembers: async (): Promise<Member[]> => {
        const response = await axios.get(`${API_URL}/Member`);
        return response.data;
    },
    getMemberById: async (id: number): Promise<Member> => {
        const response = await axios.get(`${API_URL}/Member/${id}`);
        return response.data;
    },
    createMember: async (member: Member): Promise<Member> => {
        const response = await axios.post(`${API_URL}/Member`, member);
        return response.data;
    },
};


// Booking API Service
export const BookingService = {
    getBookings: async (): Promise<Booking[]> => {
        const response = await axios.get(`${API_URL}/bookings`);
        return response.data;
    },
    getBookingById: async (id: number): Promise<Booking> => {
        const response = await axios.get(`${API_URL}/bookings/${id}`);
        return response.data;
    },
    createBooking: async (booking: Booking): Promise<Booking> => {
        const response = await axios.post(`${API_URL}/bookings`, booking);
        return response.data;
    },
};


// Benefits API Service
export const BenefitsService = {
    getBenefits: async (): Promise<Benefits[]> => {
        const response = await axios.get(`${API_URL}/Benefits`);
        return response.data;
    },
    getBenefitsById: async (id: number): Promise<Benefits> => {
        const response = await axios.get(`${API_URL}/Benefits/${id}`);
        return response.data;
    },
    createBenefits: async (benefit: Benefits): Promise<Benefits> => {
        const response = await axios.post(`${API_URL}/Benefits`, benefit);
        return response.data;
    },
};

// Payment API Service
export const PaymentService = {
    getPayments: async (): Promise<Payment[]> => {
        const response = await axios.get(`${API_URL}/Payment`);
        return response.data;
    },
    getPaymentById: async (id: number): Promise<Payment> => {
        const response = await axios.get(`${API_URL}/Payment/${id}`);
        return response.data;
    },
    createPayment: async (payment: Payment): Promise<Payment> => {
        console.log("Sending Payment:", payment);
        const response = await axios.post(`${API_URL}/Payment`, payment);
        return response.data;
    },

    
};




// Mock payment function
export const mockPayment = async (paymentData: Payment): Promise<Payment> => {
    // จำลองการชำระเงิน
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve({
                ...paymentData,
                PaymentStatus: true, // ตั้งค่า PaymentStatus เป็น true
            });
        }, 1000); // ล่าช้า 1 วินาทีเพื่อจำลองการทำงาน
    });
};
