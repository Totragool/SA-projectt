export interface Payment {
    id: number;
    PaymentStatus: boolean; // true สำหรับ paid, false สำหรับ unpaid
    PaymentDate: Date;
    PaymentTime: string;
    MemberID: number;
    BookingID: number;
    BenefitID: number;
}
