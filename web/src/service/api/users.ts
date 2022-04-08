import Restful from '@/service/api/restful';

class ApiUser extends Restful {
    version = 'v1';

    resource = 'users';
}

export default new ApiUser();
