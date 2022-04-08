declare namespace ApiUser {
    interface Model {
        id: number;
        email: string;
        username: string;
        nickname: string;
        password: string;
        is_active: boolean;
        created_at: string;
        updated_at: string;
    }
}
