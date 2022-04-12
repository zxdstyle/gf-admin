import Restful from '@/service/api/restful';

class ApiMenu extends Restful {
    version = 'v1';

    resource = 'menus';
}

export default new ApiMenu();
