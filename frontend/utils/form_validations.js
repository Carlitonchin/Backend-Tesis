export const validate_email = (value)=>{
    if(!value || value.trim() === '')
        return "Debe ingresar su correo"

    var filter = /^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    if (!filter.test(value))
        return "Ingrese un correo válido"
    
    return true
}

export const validate_username = (value)=>{
    if(!value || value.trim() === '')
        return "Ingrese un nombre de usuario"
    
    return true
}

export const validate_pass = (value)=>{
    if(!value || value === '')
        return "Ingrese una contraseña"
    
    if(value.length < 6)
        return "La contraseña debe tener al menos 6 caracteres"
    
    if(value.length > 30)
        return "La contraseña debe tener máximo 30 caracteres"
    
    return true
}

export const validate_question = (value) =>{
    if(!value || value.trim() === '')
        return "Escribe una pregunta"
    
    return true
}