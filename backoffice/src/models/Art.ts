export type Art = {
    _id: string;
    show: boolean;
    order: number;
    title: string;
    category: string;
    description: string;
    date: string;
    images: Image[];
    principal: string;
}

export type Image = {
    order: number;
    url: string;
}

export type ArtResume = {
    _id: string;
    show: boolean;
    order: number;
    title: string;
    category: string;
    description: string;
    date: string;
    image: string;
}

export interface ClassRoom {
    id: string;
    name: string;
    code: string;
    owner: string;
    created_at: string;
    users: string[];
    totalScores: number;
    totalQuestions: number;
}
