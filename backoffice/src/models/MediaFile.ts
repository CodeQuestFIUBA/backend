export interface FileProps {
    position: number;
    id: string;
    file: File;
    uploading: boolean;
    principal: boolean;
    error: boolean;
    url?: string;
}
