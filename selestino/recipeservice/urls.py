from django.urls import path
from . import views
import logging
from django.conf import settings  
from django.conf.urls.static import static
from django.contrib import admin
import time

logger = logging.getLogger(__name__)

def log_request(request, view_name):
    logger.debug(f"Request received for view: {view_name}, Method: {request.method}, Path: {request.path}, Time: {time.ctime()}")

urlpatterns = [
    path('', views.recipe_list, name='recipe_list'),  
    path('recipe/<int:id>/', views.recipe_detail, name='recipe_detail'),
    path('recipe/add/', views.add_recipe, name='add_recipe'),
    path('recipe/<int:recipe_id>/review/', views.add_review, name='add_review'),

]

if settings.DEBUG:
    urlpatterns += static(settings.STATIC_URL, document_root=settings.STATIC_ROOT)
