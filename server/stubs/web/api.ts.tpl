import Restful from '@/service/api/restful';

class Api${ .name } extends Restful {
    version = '${ .version }';

    resource = '${ .resource }';
}

export default new Api${ .name }();
