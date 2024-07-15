export type EditArt = {
    title: string;
    category: string;
    description?: string;
    date?: string;
    images: {
        _id: string;
        id: string;
        order: number;
        principal: boolean;
        url: string;
    }[];
}
