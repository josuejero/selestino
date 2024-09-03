from django.urls import path
from . import views


# urls.py

from django.conf import settings
from django.conf.urls.static import static
from django.urls import include
from django.contrib import admin

urlpatterns = [
    path('', views.recipe_list, name='recipe_list'),
    path('recipe/<int:id>/', views.recipe_detail, name='recipe_detail'),
    path('recipe/add/', views.add_recipe, name='add_recipe'),
]

if settings.DEBUG:
    urlpatterns += static(settings.STATIC_URL, document_root=settings.STATIC_ROOT)