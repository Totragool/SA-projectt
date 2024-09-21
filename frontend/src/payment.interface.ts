export interface Payment {
    id: number;
    PaymentStatus: boolean;
    PaymentDate: Date;
    PaymentTime: string;
    MemberID: number;
    BookingID: number;
    BenefitID: number;
}
