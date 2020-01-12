export function objectToFormData(object) {
    const formData = new FormData();
    Object.keys(object).forEach((key) => {
        formData.append(key, object[key] instanceof Date ? object[key].toISOString() : object[key]);
    });
    return formData;
}
