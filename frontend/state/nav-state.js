export default (current)=>
{
    let result =  [
            {id:0,text:"Inicio", link:"/", current:false},
            {id:1,text:"FAQ", link:"/faq", current:false},
            {id:2,text:"Dashbord", link:"/dashbord", current:false}
        ]

    switch(current)
    {
        case "Inicio": result[0].current = true; break;

        case "FAQ": result[1].current = true; break;

        case "Dashbord": result[2].current = true; break;
    }

    return result;
}