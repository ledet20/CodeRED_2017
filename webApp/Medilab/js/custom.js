(function ($) {
    
    // Add smooth scrolling to all links in navbar
    $(".navbar a,a.btn-appoint, .quick-info li a, .overlay-detail a").on('click', function(event) {
        
        var hash = this.hash;
        if( hash ) {
            event.preventDefault();
            $('html, body').animate({
              scrollTop: $(hash).offset().top
          }, 900, function(){
              window.location.hash = hash;
          });
        }

    });
       
    //jQuery to collapse the navbar on scroll
    $(window).scroll(function() {
        if ($(".navbar-default").offset().top > 50) {
            $(".navbar-fixed-top").addClass("top-nav-collapse");
        } else {
            $(".navbar-fixed-top").removeClass("top-nav-collapse");
        }
    });
    
})(jQuery);


$(document).ready(function(){ 
	alert("work " + ticketNum);
	
	var ticketNum = window.$vars.value_out;

    for(var i = 0; i<ticketNum;i++){
     var value = (i + 1);
   
	$('<div id="addedDiv_num" />').text(value + ':').appendTo('.containerDischarge');
	 //$('<div/>').appendTo('.containerDischarge').text("here is some text");   
       }
});


$('.button').click(function() {
	client.publish("distrk/nfc-swipe",$('.input').val());
	//alert($('.input').val());
	//alert($(this).value);
	//client.publish("distrk/nfc-swipe", "1");
});











