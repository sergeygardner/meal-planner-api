package handler

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/go-chi/jwtauth/v5"
	"github.com/pkg/errors"
	DomainResponse "github.com/sergeygardner/meal-planner-api/domain/response"
	ApplicationMiddleware "github.com/sergeygardner/meal-planner-api/infrastructure/service/jwt"
	"golang.org/x/exp/slices"
	"reflect"
	"strings"
)

var (
	jwtKey         = ApplicationMiddleware.GetJwtKey()
	jwtAuth        = jwtauth.New("HS256", jwtKey, nil)
	helperCommands string
	authToken      *DomainResponse.AuthToken
	parentCommand  *string
	commands       = commandList{
		commands: map[string]commandStruct{
			"AuthCredentials": {
				Description: "the AuthCredentials command to start authentication.",
				Function:    authCredentials,
			},
			"AuthConfirmation": {
				Description: "the AuthConfirmation command to confirm authentication.",
				Function:    authConfirmation,
			},
			"AuthRefresh": {
				Description: "the AuthRefresh command to refresh authentication.",
				Function:    authRefresh,
			},
			"AuthRegister": {
				Description: "the AuthRegister command to register a user.",
				Function:    authRegister,
			},
			"GetParentIds": {
				Description: "the GetParentIds command to get any amount of parent id. see commands (SetParentId, RemoveParentId, ClearParentIds)",
				Function:    getParentIds,
			},
			"SetParentId": {
				Description: "the SetParentId command to set any amount of parent id. try one of these - (picture_id, category_id, ingredient_id, process_id, recipe_id).",
				Function:    setParentId,
			},
			"RemoveParentId": {
				Description: "the RemoveParentId command to remove any amount of parent id. try one of these - (picture_id, category_id, ingredient_id, process_id, recipe_id).",
				Function:    removeParentId,
			},
			"ClearParentIds": {
				Description: "the ClearParentId command to remove all of parent ids.",
				Function:    clearParentIds,
			},
			"AltNamesInfo": {
				Description: "the AltNamesInfo command to show all of alt names for specific parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    altNamesInfo,
			},
			"AltNameCreate": {
				Description: "the AltNameCreate command to create an alt name and show one for specific parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    altNameCreate,
			},
			"AltNameInfo": {
				Description: "the AltNameInfo command to show an alt name for specific id and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    altNameInfo,
			},
			"AltNameUpdate": {
				Description: "the AltNameUpdate command to update an alt name and show one for specific id and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    altNameUpdate,
			},
			"AltNameDelete": {
				Description: "the AltNameDelete command to delete an alt name for specific id and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    altNameDelete,
			},
			"CategoriesInfo": {
				Description: "the CategoriesInfo command to show all of categories for specific parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    categoriesInfo,
			},
			"CategoryCreate": {
				Description: "the CategoryCreate command to create a category and show one for specific parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    categoryCreate,
			},
			"CategoryInfo": {
				Description: "the CategoryInfo command to show a category for specific id and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    categoryInfo,
			},
			"CategoryUpdate": {
				Description: "the CategoryUpdate command to update a category and show one for specific id and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    categoryUpdate,
			},
			"CategoryDelete": {
				Description: "the CategoryDelete command to delete a category for specific id and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds).",
				Function:    categoryDelete,
			},
			"IngredientsInfo": {
				Description: "the IngredientsInfo command to show all of ingredients for specific user.",
				Function:    ingredientsInfo,
			},
			"IngredientCreate": {
				Description: "the IngredientCreate command to create an ingredient and show one for specific user.",
				Function:    ingredientCreate,
			},
			"IngredientInfo": {
				Description: "the IngredientInfo command to show an ingredient for specific id and user.",
				Function:    ingredientInfo,
			},
			"IngredientUpdate": {
				Description: "the IngredientUpdate command to update an ingredient and show one for specific id and user.",
				Function:    ingredientUpdate,
			},
			"IngredientDelete": {
				Description: "the IngredientDelete command to delete an ingredient for specific id and user.",
				Function:    ingredientDelete,
			},
			"PicturesInfo": {
				Description: "the PicturesInfo command to show all of pictures for specific user.",
				Function:    picturesInfo,
			},
			"PictureCreate": {
				Description: "the PictureCreate command to create a picture and show one for specific user.",
				Function:    pictureCreate,
			},
			"PictureInfo": {
				Description: "the PictureInfo command to show a picture for specific id and user.",
				Function:    pictureInfo,
			},
			"PictureUpdate": {
				Description: "the PictureUpdate command to update a picture and show one for specific id and user.",
				Function:    pictureUpdate,
			},
			"PictureDelete": {
				Description: "the PictureDelete command to delete a picture for specific id and user.",
				Function:    pictureDelete,
			},
			"PlannersInfo": {
				Description: "the PlannersInfo command to show all of planners for specific user.",
				Function:    plannersInfo,
			},
			"PlannerCreate": {
				Description: "the PlannerCreate command to create a planner and show one for specific user.",
				Function:    plannerCreate,
			},
			"PlannerInfo": {
				Description: "the PlannerInfo command to show a planner for specific id and user.",
				Function:    plannerInfo,
			},
			"PlannerUpdate": {
				Description: "the PlannerUpdate command to update a planner and show one for specific id and user.",
				Function:    plannerUpdate,
			},
			"PlannerDelete": {
				Description: "the PlannerDelete command to delete a planner for specific id and user.",
				Function:    plannerDelete,
			},
			"PlannerIntervalsInfo": {
				Description: "the PlannerIntervalsInfo command to show all of planner intervals for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerIntervalsInfo,
			},
			"PlannerIntervalCreate": {
				Description: "the PlannerIntervalCreate command to create a planner interval and show one for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerIntervalCreate,
			},
			"PlannerIntervalInfo": {
				Description: "the PlannerIntervalInfo command to show a planner interval for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerIntervalInfo,
			},
			"PlannerIntervalUpdate": {
				Description: "the PlannerIntervalUpdate command to update a planner interval and show one for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerIntervalUpdate,
			},
			"PlannerIntervalDelete": {
				Description: "the PlannerIntervalDelete command to delete a planner interval for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerIntervalDelete,
			},
			"PlannerRecipesInfo": {
				Description: "the PlannerRecipesInfo command to show all of planner recipes for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerRecipesInfo,
			},
			"PlannerRecipeCreate": {
				Description: "the PlannerRecipeCreate command to create a planner recipe and show one for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerRecipeCreate,
			},
			"PlannerRecipeInfo": {
				Description: "the PlannerRecipeInfo command to show a planner recipe for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerRecipeInfo,
			},
			"PlannerRecipeUpdate": {
				Description: "the PlannerRecipeUpdate command to update a planner recipe and show one for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerRecipeUpdate,
			},
			"PlannerRecipeDelete": {
				Description: "the PlannerRecipeDelete command to delete a planner recipe for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    plannerRecipeDelete,
			},
			"RecipesInfo": {
				Description: "the RecipesInfo command to show all of recipes for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipesInfo,
			},
			"RecipeCreate": {
				Description: "the RecipeCreate command to create a recipe and show one for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeCreate,
			},
			"RecipeInfo": {
				Description: "the RecipeInfo command to show a recipe for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeInfo,
			},
			"RecipeUpdate": {
				Description: "the RecipeUpdate command to update a recipe and show one for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeUpdate,
			},
			"RecipeDelete": {
				Description: "the RecipeDelete command to delete a recipe for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeDelete,
			},
			"RecipeCategoriesInfo": {
				Description: "the RecipeCategoriesInfo command to show all of recipe categories for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeCategoriesInfo,
			},
			"RecipeCategoryCreate": {
				Description: "the RecipeCategoryCreate command to create a recipe category and show one for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeCategoryCreate,
			},
			"RecipeCategoryInfo": {
				Description: "the RecipeCategoryInfo command to show a recipe category for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeCategoryInfo,
			},
			"RecipeCategoryUpdate": {
				Description: "the RecipeCategoryUpdate command to update a recipe category and show one for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeCategoryUpdate,
			},
			"RecipeCategoryDelete": {
				Description: "the RecipeCategoryDelete command to delete a recipe category for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeCategoryDelete,
			},
			"RecipeIngredientsInfo": {
				Description: "the RecipeIngredientsInfo command to show all of recipe ingredients for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeIngredientsInfo,
			},
			"RecipeIngredientCreate": {
				Description: "the RecipeIngredientCreate command to create a recipe ingredient and show one for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeIngredientCreate,
			},
			"RecipeIngredientInfo": {
				Description: "the RecipeIngredientInfo command to show a recipe ingredient for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeIngredientInfo,
			},
			"RecipeIngredientUpdate": {
				Description: "the RecipeIngredientUpdate command to update a recipe ingredient and show one for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeIngredientUpdate,
			},
			"RecipeIngredientDelete": {
				Description: "the RecipeIngredientDelete command to delete a recipe ingredient for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeIngredientDelete,
			},
			"RecipeMeasuresInfo": {
				Description: "the RecipeMeasuresInfo command to show all of recipe measures for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeMeasuresInfo,
			},
			"RecipeMeasureCreate": {
				Description: "the RecipeMeasureCreate command to create a recipe measure and show one for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeMeasureCreate,
			},
			"RecipeMeasureInfo": {
				Description: "the RecipeMeasureInfo command to show a recipe measure for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeMeasureInfo,
			},
			"RecipeMeasureUpdate": {
				Description: "the RecipeMeasureUpdate command to update a recipe measure and show one for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeMeasureUpdate,
			},
			"RecipeMeasureDelete": {
				Description: "the RecipeMeasureDelete command to delete a recipe measure for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeMeasureDelete,
			},
			"RecipeProcessesInfo": {
				Description: "the RecipeProcessesInfo command to show all of recipe processes for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeProcessesInfo,
			},
			"RecipeProcessCreate": {
				Description: "the RecipeProcessCreate command to create a recipe process and show one for specific user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeProcessCreate,
			},
			"RecipeProcessInfo": {
				Description: "the RecipeProcessInfo command to show a recipe process for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeProcessInfo,
			},
			"RecipeProcessUpdate": {
				Description: "the RecipeProcessUpdate command to update a recipe process and show one for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeProcessUpdate,
			},
			"RecipeProcessDelete": {
				Description: "the RecipeProcessDelete command to delete a recipe process for specific id and user and parent_id. see commands (SetParentId, RemoveParentId, ClearParentIds)..",
				Function:    recipeProcessDelete,
			},
			"UnitsInfo": {
				Description: "the UnitsInfo command to show all of units for specific user.",
				Function:    unitsInfo,
			},
			"UnitCreate": {
				Description: "the UnitCreate command to create a unit and show one for specific user.",
				Function:    unitCreate,
			},
			"UnitInfo": {
				Description: "the UnitInfo command to show a unit for specific id and user.",
				Function:    unitInfo,
			},
			"UnitUpdate": {
				Description: "the UnitUpdate command to update a unit and show one for specific id and user.",
				Function:    unitUpdate,
			},
			"UnitDelete": {
				Description: "the UnitDelete command to delete a unit for specific id and user.",
				Function:    unitDelete,
			},
			"UserInfo": {
				Description: "the UserInfo command to show a user for specific id and user.",
				Function:    userInfo,
			},
			"UserUpdate": {
				Description: "the UserUpdate command to update a user and show one for specific id and user.",
				Function:    userUpdate,
			},
			"UserDelete": {
				Description: "the UserDelete command to delete a user for specific id and user.",
				Function:    userDelete,
			},
			"auth": {
				Description: "the auth command to Help faster authentication for username=username.",
				Function:    auth,
			},
			"Help": {
				Description: "the Help command to show the Help message.",
				Function:    Help,
			},
			"exit": {
				Description: "the exit command from the application.",
				Function:    exit,
			},
		},
	}
	errorCommandCalling        = errors.New("an error occurred while running command. Wrong command name.")
	errorRemoveParentId        = errors.New("an error occurred while running command. Wrong parent_id key to remove.")
	errorSetParentIdWrong      = errors.New("an error occurred while running command. Wrong parent_id key and value to set.")
	errorSetParentIdWrongKey   = errors.New("an error occurred while running command. Wrong parent_id key to set.")
	errorSetParentIdWrongValue = errors.New("an error occurred while running command. Wrong parent_id value to set.")
	parentIdKeys               []string
	parentIdValues             []string
)

type commandList struct {
	commands map[string]commandStruct
}

type commandStruct struct {
	Description string
	Function    func(command string) (int, error)
}

func init() {
	helperCommands = commands.HelpMessage()
}

func (cl *commandList) Call(commandName string, message string) (int, error) {
	commandStructItem, ok := cl.commands[commandName]

	if !ok {
		return resetParentCommand(StatusNotFound, errorCommandCalling)
	}

	statusCalling, errorCalling := commandStructItem.Function(message)

	return resetParentCommand(statusCalling, errorCalling)
}

func (cl *commandList) HelpMessage() string {
	var response []string

	reflectCommands := reflect.ValueOf(cl.commands)

	for _, key := range reflectCommands.MapKeys() {
		commandStructItem := reflectCommands.MapIndex(key).Interface().(commandStruct)

		response = append(response, commandStructItem.Description)
	}

	slices.Sort(response)

	return strings.Join(response, "\n")
}

const (
	StatusOk = iota
	StatusExit
	StatusError
	StatusContinue
	StatusNotFound
)

func Run(command string) (int, error) {
	if parentCommand == nil {
		parentCommand = &command
	}

	statusCalling, errorCalling := commands.Call(*parentCommand, command)

	if statusCalling == StatusNotFound {
		showErrorMessage("the command '%s' is not found. try these commands:\n\n%s\n", command, helperCommands)
		return StatusError, errorCalling
	} else if errorCalling != nil {
		return StatusError, errorCalling
	}

	return statusCalling, nil
}

func getParentIds(_ string) (int, error) {
	showInfoMessage("parent id keys are %v, parent id values are %v", parentIdKeys, parentIdValues)

	return StatusOk, nil
}

func setParentId(message string) (int, error) {
	if message == "SetParentId" {
		showDialogMessage("input parent_id key and value and separate them with space e.g. recipe_id 00000000-0000-0000-0000-000000000000")
		return StatusContinue, nil
	}

	pair := strings.Split(message, " ")

	if len(pair) != 2 {
		return StatusError, errorSetParentIdWrong
	}

	key := pair[0]

	if key == "" {
		return StatusError, errorSetParentIdWrongKey
	}
	value := pair[1]

	if value == "" {
		return StatusError, errorSetParentIdWrongValue
	}

	parentIdKeys = append(parentIdKeys, pair[0])
	parentIdValues = append(parentIdValues, pair[1])

	return resetParentCommand(StatusOk, nil)
}

func removeParentId(message string) (int, error) {
	var (
		i     int
		key   string
		value string
	)

	if message == "RemoveParentId" {
		return StatusContinue, nil
	}

	for i, key = range parentIdKeys {
		if key == message {
			value = parentIdValues[i]

			parentIdKeys = slices.Delete(parentIdKeys, i, i)
			parentIdValues = slices.Delete(parentIdValues, i, i)
		}
	}

	if key == "" || value == "" {
		return StatusError, errorRemoveParentId
	}

	return StatusOk, nil
}

func clearParentIds(_ string) (int, error) {
	parentIdKeys = []string{}
	parentIdValues = []string{}

	return StatusOk, nil
}

func Help(_ string) (int, error) {
	showInfoMessage("the Help message. to interact with the application try these commands:\n\n%s\n", helperCommands)

	return StatusOk, nil
}

func exit(_ string) (int, error) {
	showInfoMessage("bye-bye")

	return StatusExit, nil
}

func showErrorMessage(message string, arguments ...any) {
	if arguments == nil {
		fmt.Println(color.Ize(color.Red, message))
	} else {
		fmt.Println(color.Ize(color.Red, fmt.Sprintf(message, arguments...)))
	}
}

func showInfoMessage(message string, arguments ...any) {
	if arguments == nil {
		fmt.Println(color.Ize(color.Yellow, message))
	} else {
		fmt.Println(color.Ize(color.Yellow, fmt.Sprintf(message, arguments...)))
	}
}

func showDialogMessage(message string, arguments ...any) {
	if arguments == nil {
		fmt.Println(color.Ize(color.Gray, message))
	} else {
		fmt.Println(color.Ize(color.Gray, fmt.Sprintf(message, arguments...)))
	}
}

func resetParentCommand(statusCommand int, errorCommand error) (int, error) {
	switch statusCommand {
	case StatusOk:
		parentCommand = nil
	case StatusError:
		parentCommand = nil
	case StatusNotFound:
		parentCommand = nil
	case StatusContinue:
	default:
		parentCommand = nil
	}

	return statusCommand, errorCommand
}
