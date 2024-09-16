from rest_framework.decorators import api_view, permission_classes
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response
from rest_framework import status
from .models import Recipe
from .serializers import RecipeSerializer
from django.shortcuts import get_object_or_404
import logging

logger = logging.getLogger(__name__)

@api_view(['GET'])
def recipe_list(request):
    logger.debug(f"Accessing recipe_list. Method: {request.method}, Path: {request.path}")
    recipes = Recipe.objects.all()
    serializer = RecipeSerializer(recipes, many=True)
    return Response(serializer.data)

@api_view(['GET', 'PUT', 'DELETE'])
def recipe_detail(request, id):
    recipe = get_object_or_404(Recipe, id=id)

    if request.method == 'GET':
        serializer = RecipeSerializer(recipe)
        return Response(serializer.data)

    elif request.method == 'PUT':
        serializer = RecipeSerializer(recipe, data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

    elif request.method == 'DELETE':
        recipe.delete()
        return Response(status=status.HTTP_204_NO_CONTENT)

@api_view(['POST'])
@permission_classes([IsAuthenticated])
def add_recipe(request):
    if request.method == 'POST':
        serializer = RecipeSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

@api_view(['POST'])
@permission_classes([IsAuthenticated])
def delete_recipe(request, id):
    recipe = get_object_or_404(Recipe, id=id)
    if request.user.has_perm('recipeservice.delete_recipe'):
        recipe.delete()
        return Response(status=status.HTTP_204_NO_CONTENT)
    return Response({'error': 'Unauthorized'}, status=status.HTTP_403_FORBIDDEN)

@api_view(['PUT'])
@permission_classes([IsAuthenticated])
def edit_recipe(request, id):
    recipe = get_object_or_404(Recipe, id=id)
    if request.method == 'PUT':
        serializer = RecipeSerializer(recipe, data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

@api_view(['POST'])
@permission_classes([IsAuthenticated])
def add_review(request, recipe_id):
    recipe = get_object_or_404(Recipe, id=recipe_id)
    if request.method == 'POST':
        
        serializer = ReviewSerializer(data=request.data)
        if serializer.is_valid():
            review = serializer.save(commit=False)
            review.recipe = recipe
            review.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
