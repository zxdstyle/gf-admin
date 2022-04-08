declare namespace ApiScaffold {
    interface Table {
        name: string;
        engine: string;
    }

    interface column {
        name: string;
        label: string;
        type: string;
        length: number;
        nullable: boolean;
        default_value: string | number;
        comment: string;
    }

    interface Field {
        index: number;
        name: string;
        json_name: string;
        translation: string;
        type: string;
        nullable: boolean;
        default_value: string;
        comment: string;
        hidden: boolean;
    }

    interface FormModel {
        table: string;
        model: string;
        handler: string;
        view_path: string;
        repository: string;
        create_view: boolean;
        create_model: boolean;
        create_repository: boolean;
        create_handler: boolean;
        create_lang: boolean;
        fields: Field[];
    }

    interface Preview {
        lang?: string;
        path?: string;
        content?: string;
        children?: Preview[];
    }
}
