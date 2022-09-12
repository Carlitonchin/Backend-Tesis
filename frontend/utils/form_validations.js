export const validate_email = (value)=>{
    if(!value || value === '')
        return "Debe ingresar su correo"

    var filter = /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    if (!filter.test(value))
        return "Ingrese un correo válido"
    
    return true
}