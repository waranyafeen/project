import { Role } from "@src/providers/auth";

interface Menu {
    label: string;
    to: string;
}
type TMenus = {
    [key in Role]: Menu[];
};

export const Menus: TMenus = {
    employee: [
        {
            label: "Ticket",
            to: "/ticket",
        },
    ],
    admin: [
        {
            label: "Ticket",
            to: "/ticket",
        },
    ],
    user: [
        {
            label: "Ticket",
            to: "/ticket",
        },
    ]

};
